package commands

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"buybikeshop/libs/go/ngutils/templates"

	"github.com/spf13/cobra"
)

// Usage:
// ngutils create --entities-path=./apps/office/client/src/app/entities --entity-name=partner --service=office
// --entities-path - path for angular ngrx entities folder
// --entity-name   - name of the entity to create (singular, e.g. "partner")
// --service       - service name for API routes (auto-derived from entities-path if omitted)

func NewNGRXEntityCommand() (cmd *cobra.Command) {
	var entitiesPath string
	var entityName string
	var service string

	cmd = &cobra.Command{
		Use:   "create",
		Short: "Create new NGRX entity",
		Args:  cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if entitiesPath == "" {
				return errors.New("--entities-path is required")
			}
			if entityName == "" {
				return errors.New("--entity-name is required")
			}

			if service == "" {
				service = deriveService(entitiesPath)
			}
			if service == "" {
				return errors.New("--service is required (or use an entities-path like apps/<service>/client/...)")
			}

			return generateNGRXEntity(entitiesPath, entityName, service)
		},
	}

	cmd.Flags().StringVar(&entitiesPath, "entities-path", "", "Path to Angular NgRx entities folder")
	cmd.Flags().StringVar(&entityName, "entity-name", "", "Name of the entity to create (singular, e.g. 'partner')")
	cmd.Flags().StringVar(&service, "service", "", "Service name for API routes (e.g. 'office')")

	return cmd
}

/*

result structure:
- ./<entity_name>/
-	./index.ts
-	./api
-	-	./api.service.ts
-	./model
-	-	./index.ts
-	-	./<entity_name>.actions.ts
-	-	./<entity_name>.effects.ts
-	-	./<entity_name>.model.ts
-	-	./<entity_name>.reducer.ts
-	-	./<entity_name>.selectors.ts
-	./ui/
-	-	./index.ts
*/

type entityData struct {
	EntityName  string // PascalCase singular: "Partner"
	EntityModel string // PascalCase + Model: "PartnerModel"
	Module      string // lowercase singular: "partner"
	Service     string // service name: "office"
}

// deriveService extracts the service name from an entities path like "apps/<service>/client/..."
func deriveService(entitiesPath string) string {
	parts := strings.Split(filepath.ToSlash(entitiesPath), "/")
	for i, part := range parts {
		if part == "apps" && i+1 < len(parts) {
			return parts[i+1]
		}
	}
	return ""
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func generateNGRXEntity(entitiesPath, entityName, service string) error {
	module := strings.ToLower(entityName)
	data := entityData{
		EntityName:  capitalize(module),
		EntityModel: capitalize(module) + "Model",
		Module:      module,
		Service:     service,
	}

	basePath := filepath.Join(entitiesPath, module)

	// Create directory structure
	dirs := []string{
		filepath.Join(basePath, "api"),
		filepath.Join(basePath, "model"),
		filepath.Join(basePath, "ui"),
	}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("creating directory %s: %w", dir, err)
		}
	}

	// Generate files from templates
	tmplFiles := []struct {
		template string
		output   string
	}{
		{"ngrx_entity/api.service.tmpl", filepath.Join("api", "api.service.ts")},
		{"ngrx_entity/index.tmpl", filepath.Join("model", "index.ts")},
		{"ngrx_entity/product.actions.tmpl", filepath.Join("model", module+".actions.ts")},
		{"ngrx_entity/product.effects.tmpl", filepath.Join("model", module+".effects.ts")},
		{"ngrx_entity/product.model.tmpl", filepath.Join("model", module+".model.ts")},
		{"ngrx_entity/product.reducer.tmpl", filepath.Join("model", module+".reducer.ts")},
		{"ngrx_entity/product.selectors.tmpl", filepath.Join("model", module+".selectors.ts")},
	}

	for _, f := range tmplFiles {
		outPath := filepath.Join(basePath, f.output)
		if err := generateFile(f.template, outPath, data); err != nil {
			return fmt.Errorf("generating %s: %w", outPath, err)
		}
		fmt.Printf("  created: %s\n", outPath)
	}

	// Root index.ts - barrel export
	indexPath := filepath.Join(basePath, "index.ts")
	if err := os.WriteFile(indexPath, []byte("export * from './model';\n"), 0644); err != nil {
		return fmt.Errorf("writing %s: %w", indexPath, err)
	}
	fmt.Printf("  created: %s\n", indexPath)

	// Empty ui/index.ts placeholder
	uiIndexPath := filepath.Join(basePath, "ui", "index.ts")
	if err := os.WriteFile(uiIndexPath, []byte(""), 0644); err != nil {
		return fmt.Errorf("writing %s: %w", uiIndexPath, err)
	}
	fmt.Printf("  created: %s\n", uiIndexPath)

	fmt.Printf("\nEntity '%s' generated at %s\n", module, basePath)
	return nil
}

func generateFile(tmplPath, outPath string, data entityData) error {
	content, err := templates.NgrxEntity.ReadFile(tmplPath)
	if err != nil {
		return fmt.Errorf("reading template %s: %w", tmplPath, err)
	}

	tmpl, err := template.New(filepath.Base(tmplPath)).Parse(string(content))
	if err != nil {
		return fmt.Errorf("parsing template %s: %w", tmplPath, err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("executing template %s: %w", tmplPath, err)
	}

	return os.WriteFile(outPath, buf.Bytes(), 0644)
}

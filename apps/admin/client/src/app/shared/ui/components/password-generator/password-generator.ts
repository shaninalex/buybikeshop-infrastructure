import { Component, EventEmitter, Output } from '@angular/core';

@Component({
    selector: 'app-password-generator',
    imports: [],
    template: `<button class="btn btn-sm btn-outline-info" type="button" (click)="generatePassword()">
        Generate
    </button>`,
})
export class PasswordGenerator {
    @Output() password = new EventEmitter();

    generatePassword(): void {
        this.password.emit(this.generateSecurePassword());
    }

    private generateSecurePassword(length: number = 12): string {
        const uppercase = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
        const lowercase = 'abcdefghijklmnopqrstuvwxyz';
        const numbers = '0123456789';
        const symbols = '!@#$%^&*()-_=+[]{}|;:,.<>?';

        const allChars = uppercase + lowercase + numbers + symbols;
        const requiredChars = [
            this.getRandomChar(uppercase),
            this.getRandomChar(lowercase),
            this.getRandomChar(numbers),
            this.getRandomChar(symbols),
        ];

        const remainingLength = length - requiredChars.length;
        const passwordChars = [...requiredChars];

        const randomValues = new Uint32Array(remainingLength);
        crypto.getRandomValues(randomValues);

        for (let i = 0; i < remainingLength; i++) {
            passwordChars.push(allChars[randomValues[i] % allChars.length]);
        }
        return this.shuffle(passwordChars).join('');
    }

    private getRandomChar(charset: string): string {
        const randomValue = new Uint32Array(1);
        crypto.getRandomValues(randomValue);
        return charset[randomValue[0] % charset.length];
    }

    private shuffle(array: string[]): string[] {
        const randomValues = new Uint32Array(array.length);
        crypto.getRandomValues(randomValues);

        for (let i = array.length - 1; i > 0; i--) {
            const j = randomValues[i] % (i + 1);
            [array[i], array[j]] = [array[j], array[i]];
        }
        return array;
    }
}

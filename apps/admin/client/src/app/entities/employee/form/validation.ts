import { email, PathKind, pattern, required, SchemaPathTree, validate } from '@angular/forms/signals';
import { EmployeeCreateFormModel } from '@entities/employee/model/employee.model';
import { isAfter } from 'date-fns';

export function employeeFormValidation(schemaPath: SchemaPathTree<EmployeeCreateFormModel, PathKind.Root>) {
    email(schemaPath.email, {message: 'Invalid email address format'});
    required(schemaPath.email, {message: 'Email is required'});

    required(schemaPath.name, {message: 'Name is required'});
    pattern(
        schemaPath.name,
        /^[\p{L} ]+$/u,
        { message: 'Name may contain only letters and spaces' }
    );

    validate(schemaPath.dob, ({value}) => {
        if (isAfter(value(), new Date())) {
            return {
                kind: 'dob',
                message: 'Date of birth can`t be future date',
            };
        }
        return null;
    });

    // Phone validator: digits, +, (), space, and dash
    pattern(
        schemaPath.phone,
        /^$|^[0-9+()\- ]+$/,
        {
            message: 'Phone number may contain only digits, spaces, "+", "-", and parentheses'
        }
    );

    if (schemaPath.password) {
        required(schemaPath.password, {message: 'Passwords are required'});
    }
}

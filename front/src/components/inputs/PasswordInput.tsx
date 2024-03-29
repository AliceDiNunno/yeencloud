import React, { useState, ChangeEvent } from 'react';
import { Input } from "@mui/joy";

interface CustomInputProps {
    value: string,
    onChange: (value: string) => void
}

const PasswordInput: React.FC<CustomInputProps> = ({ value, onChange }) => (
    <Input
        type="text"
        value={value}
        onChange={(event) => onChange(event.target.value)}
    />
);

export default PasswordInput;
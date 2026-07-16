import type { OperationId, OperationMeta } from '../types';

export const OPERATIONS: OperationMeta[] = [
  { id: 'add', label: 'Add', symbol: '+', unary: false },
  { id: 'subtract', label: 'Subtract', symbol: '−', unary: false },
  { id: 'multiply', label: 'Multiply', symbol: '×', unary: false },
  { id: 'divide', label: 'Divide', symbol: '÷', unary: false },
  { id: 'power', label: 'Power', symbol: '^', unary: false },
  { id: 'percentage', label: 'Percent', symbol: '%', unary: false },
  { id: 'sqrt', label: 'Square root', symbol: '√', unary: true },
];

export function getOperation(id: OperationId): OperationMeta {
  const operation = OPERATIONS.find((candidate) => candidate.id === id);
  if (!operation) throw new Error(`unknown operation: ${id}`);
  return operation;
}

export type FieldValidation = { valid: true; value: number } | { valid: false; error: string };

export function validateOperand(raw: string, fieldName: string): FieldValidation {
  const trimmed = raw.trim();
  if (trimmed === '') {
    return { valid: false, error: `${fieldName} is required` };
  }
  const value = Number(trimmed);
  if (!Number.isFinite(value)) {
    return { valid: false, error: `${fieldName} must be a valid number` };
  }
  return { valid: true, value };
}

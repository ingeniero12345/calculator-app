import type { OperationId, OperationMeta } from '../types';

/**
 * Ordered list of operations rendered in the UI. Keeping this metadata in one
 * place means adding an operation requires a single entry here plus its backend
 * counterpart.
 */
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
  const op = OPERATIONS.find((o) => o.id === id);
  if (!op) throw new Error(`unknown operation: ${id}`);
  return op;
}

/** Result of validating a single operand input field. */
export type FieldValidation = { valid: true; value: number } | { valid: false; error: string };

/**
 * Validates raw text from an input field. Empty input and anything that is not
 * a finite number is rejected with a human-readable message. Client-side
 * validation gives instant feedback; the backend re-validates authoritatively.
 */
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

import { describe, it, expect } from 'vitest';
import { OPERATIONS, getOperation, validateOperand } from './operations';

describe('getOperation', () => {
  it('returns metadata for a known operation', () => {
    expect(getOperation('add').label).toBe('Add');
    expect(getOperation('sqrt').unary).toBe(true);
  });

  it('throws for an unknown operation', () => {
    // @ts-expect-error invalid id
    expect(() => getOperation('modulo')).toThrow();
  });

  it('marks only sqrt as unary', () => {
    const unary = OPERATIONS.filter((o) => o.unary).map((o) => o.id);
    expect(unary).toEqual(['sqrt']);
  });
});

describe('validateOperand', () => {
  it('accepts integers and decimals', () => {
    expect(validateOperand('42', 'First value')).toEqual({ valid: true, value: 42 });
    expect(validateOperand('3.14', 'First value')).toEqual({ valid: true, value: 3.14 });
    expect(validateOperand('-5', 'First value')).toEqual({ valid: true, value: -5 });
  });

  it('trims surrounding whitespace', () => {
    expect(validateOperand('  7  ', 'First value')).toEqual({ valid: true, value: 7 });
  });

  it('rejects empty input', () => {
    const result = validateOperand('   ', 'First value');
    expect(result).toEqual({ valid: false, error: 'First value is required' });
  });

  it('rejects non-numeric input', () => {
    const result = validateOperand('abc', 'Second value');
    expect(result).toEqual({ valid: false, error: 'Second value must be a valid number' });
  });

  it('rejects Infinity', () => {
    expect(validateOperand('Infinity', 'First value').valid).toBe(false);
  });
});

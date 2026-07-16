import { useState, type FormEvent } from 'react';
import type { OperationId } from '../types';
import { OPERATIONS, getOperation, validateOperand } from '../lib/operations';
import { calculate, ApiError } from '../api/calculatorApi';
import './Calculator.css';

type Status =
  | { kind: 'idle' }
  | { kind: 'loading' }
  | { kind: 'success'; result: number }
  | { kind: 'error'; message: string };

const DISPLAY_PRECISION = 12;

export function Calculator() {
  const [operationId, setOperationId] = useState<OperationId>('add');
  const [a, setA] = useState('');
  const [b, setB] = useState('');
  const [fieldErrors, setFieldErrors] = useState<{ a?: string; b?: string }>({});
  const [status, setStatus] = useState<Status>({ kind: 'idle' });

  const operation = getOperation(operationId);

  async function handleSubmit(event: FormEvent) {
    event.preventDefault();

    const aResult = validateOperand(a, 'First value');
    const bResult = operation.unary ? null : validateOperand(b, 'Second value');

    const errors: { a?: string; b?: string } = {};
    if (!aResult.valid) errors.a = aResult.error;
    if (bResult && !bResult.valid) errors.b = bResult.error;
    setFieldErrors(errors);
    if (Object.keys(errors).length > 0) {
      setStatus({ kind: 'idle' });
      return;
    }

    setStatus({ kind: 'loading' });
    try {
      const payload = {
        a: (aResult as { value: number }).value,
        ...(operation.unary ? {} : { b: (bResult as { value: number }).value }),
      };
      const response = await calculate(operationId, payload);
      setStatus({ kind: 'success', result: response.result });
    } catch (error) {
      const message = error instanceof ApiError ? error.message : 'Something went wrong';
      setStatus({ kind: 'error', message });
    }
  }

  function handleOperationChange(id: OperationId) {
    setOperationId(id);
    setStatus({ kind: 'idle' });
    setFieldErrors({});
  }

  return (
    <section className="calculator" aria-labelledby="calc-title">
      <h1 id="calc-title" className="calculator__title">
        Calculator
      </h1>

      <div className="calculator__operations" role="group" aria-label="Choose an operation">
        {OPERATIONS.map((op) => (
          <button
            key={op.id}
            type="button"
            className={`op-button ${op.id === operationId ? 'op-button--active' : ''}`}
            aria-pressed={op.id === operationId}
            onClick={() => handleOperationChange(op.id)}
          >
            <span className="op-button__symbol" aria-hidden="true">
              {op.symbol}
            </span>
            {op.label}
          </button>
        ))}
      </div>

      <form className="calculator__form" onSubmit={handleSubmit} noValidate>
        <div className="field">
          <label htmlFor="operand-a">First value</label>
          <input
            id="operand-a"
            type="number"
            step="any"
            inputMode="decimal"
            value={a}
            onChange={(event) => setA(event.target.value)}
            aria-invalid={Boolean(fieldErrors.a)}
            aria-describedby={fieldErrors.a ? 'error-a' : undefined}
            placeholder="0"
          />
          {fieldErrors.a && (
            <span id="error-a" className="field__error" role="alert">
              {fieldErrors.a}
            </span>
          )}
        </div>

        {!operation.unary && (
          <div className="field">
            <label htmlFor="operand-b">
              Second value{operationId === 'percentage' ? ' (percent)' : ''}
            </label>
            <input
              id="operand-b"
              type="number"
              step="any"
              inputMode="decimal"
              value={b}
              onChange={(event) => setB(event.target.value)}
              aria-invalid={Boolean(fieldErrors.b)}
              aria-describedby={fieldErrors.b ? 'error-b' : undefined}
              placeholder="0"
            />
            {fieldErrors.b && (
              <span id="error-b" className="field__error" role="alert">
                {fieldErrors.b}
              </span>
            )}
          </div>
        )}

        <button type="submit" className="calculator__submit" disabled={status.kind === 'loading'}>
          {status.kind === 'loading' ? 'Calculating…' : 'Calculate'}
        </button>
      </form>

      <output className="calculator__output" aria-live="polite">
        {status.kind === 'success' && (
          <div className="result result--success">
            <span className="result__label">Result</span>
            <span className="result__value">{formatResult(status.result)}</span>
          </div>
        )}
        {status.kind === 'error' && (
          <div className="result result--error" role="alert">
            {status.message}
          </div>
        )}
      </output>
    </section>
  );
}

function formatResult(value: number): string {
  if (Number.isInteger(value)) return String(value);
  return String(Number(value.toPrecision(DISPLAY_PRECISION)));
}

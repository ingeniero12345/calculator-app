/** Canonical identifiers for every operation the backend exposes. */
export type OperationId =
  | 'add'
  | 'subtract'
  | 'multiply'
  | 'divide'
  | 'power'
  | 'sqrt'
  | 'percentage';

/** UI metadata describing how to render and validate a single operation. */
export interface OperationMeta {
  id: OperationId;
  label: string;
  /** Mathematical symbol shown between/around operands, e.g. "+" or "√". */
  symbol: string;
  /** Unary operations (sqrt) only take operand `a`. */
  unary: boolean;
}

/** Request body sent to the backend. `b` is omitted for unary operations. */
export interface CalculationRequest {
  a: number;
  b?: number;
}

/** Successful response returned by the backend. */
export interface CalculationResponse {
  operation: string;
  a: number;
  b?: number;
  result: number;
}

export type OperationId =
  | 'add'
  | 'subtract'
  | 'multiply'
  | 'divide'
  | 'power'
  | 'sqrt'
  | 'percentage';

export interface OperationMeta {
  id: OperationId;
  label: string;
  symbol: string;
  unary: boolean;
}

export interface CalculationRequest {
  a: number;
  b?: number;
}

export interface CalculationResponse {
  operation: string;
  a: number;
  b?: number;
  result: number;
}

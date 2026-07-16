import type { CalculationRequest, CalculationResponse, OperationId } from '../types';

const API_BASE = import.meta.env.VITE_API_BASE_URL ?? '';

export class ApiError extends Error {
  constructor(
    message: string,
    public readonly status: number,
  ) {
    super(message);
    this.name = 'ApiError';
  }
}

export async function calculate(
  operation: OperationId,
  body: CalculationRequest,
): Promise<CalculationResponse> {
  let response: Response;
  try {
    response = await fetch(`${API_BASE}/api/v1/${operation}`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(body),
    });
  } catch {
    throw new ApiError('Unable to reach the calculator service. Is the backend running?', 0);
  }

  const data = await response.json().catch(() => null);

  if (!response.ok) {
    const message =
      (data && typeof data.error === 'string' && data.error) ||
      `Request failed with status ${response.status}`;
    throw new ApiError(message, response.status);
  }

  return data as CalculationResponse;
}

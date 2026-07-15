import type { CalculationRequest, CalculationResponse, OperationId } from '../types';

/**
 * Base URL for the API. In development Vite proxies `/api` to the Go backend
 * (see vite.config.ts); in production the value can be injected at build time
 * via VITE_API_BASE_URL. Defaults to a relative path so the SPA works when
 * served behind the same origin as the API.
 */
const API_BASE = import.meta.env.VITE_API_BASE_URL ?? '';

/** Error thrown when the API returns a non-2xx response. */
export class ApiError extends Error {
  constructor(
    message: string,
    public readonly status: number,
  ) {
    super(message);
    this.name = 'ApiError';
  }
}

/**
 * Calls the backend to perform a single calculation. Throws ApiError with the
 * server-provided message on failure (e.g. division by zero), or a generic
 * network error if the request could not be completed.
 */
export async function calculate(
  operation: OperationId,
  body: CalculationRequest,
): Promise<CalculationResponse> {
  let res: Response;
  try {
    res = await fetch(`${API_BASE}/api/v1/${operation}`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(body),
    });
  } catch {
    throw new ApiError('Unable to reach the calculator service. Is the backend running?', 0);
  }

  const data = await res.json().catch(() => null);

  if (!res.ok) {
    const message =
      (data && typeof data.error === 'string' && data.error) ||
      `Request failed with status ${res.status}`;
    throw new ApiError(message, res.status);
  }

  return data as CalculationResponse;
}

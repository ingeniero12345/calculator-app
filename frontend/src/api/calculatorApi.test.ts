import { describe, it, expect, vi, afterEach } from 'vitest';
import { calculate, ApiError } from './calculatorApi';

afterEach(() => {
  vi.restoreAllMocks();
});

function mockFetch(response: unknown, ok = true, status = 200) {
  vi.stubGlobal(
    'fetch',
    vi.fn().mockResolvedValue({
      ok,
      status,
      json: () => Promise.resolve(response),
    } as Response),
  );
}

describe('calculate', () => {
  it('returns the parsed response on success', async () => {
    mockFetch({ operation: 'add', a: 2, b: 3, result: 5 });
    const result = await calculate('add', { a: 2, b: 3 });
    expect(result.result).toBe(5);
  });

  it('sends the correct URL and body', async () => {
    const fetchMock = vi.fn().mockResolvedValue({
      ok: true,
      status: 200,
      json: () => Promise.resolve({ result: 5 }),
    } as Response);
    vi.stubGlobal('fetch', fetchMock);

    await calculate('divide', { a: 10, b: 2 });

    expect(fetchMock).toHaveBeenCalledWith(
      '/api/v1/divide',
      expect.objectContaining({
        method: 'POST',
        body: JSON.stringify({ a: 10, b: 2 }),
      }),
    );
  });

  it('throws ApiError with the server message on 4xx', async () => {
    mockFetch({ error: 'division by zero is undefined' }, false, 422);
    await expect(calculate('divide', { a: 1, b: 0 })).rejects.toThrowError(
      new ApiError('division by zero is undefined', 422),
    );
  });

  it('throws a friendly ApiError when the network fails', async () => {
    vi.stubGlobal('fetch', vi.fn().mockRejectedValue(new Error('network down')));
    await expect(calculate('add', { a: 1, b: 2 })).rejects.toMatchObject({
      name: 'ApiError',
      status: 0,
    });
  });
});

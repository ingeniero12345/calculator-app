import { describe, it, expect, vi, afterEach } from 'vitest';
import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { Calculator } from './Calculator';
import * as api from '../api/calculatorApi';

afterEach(() => {
  vi.restoreAllMocks();
});

describe('Calculator', () => {
  it('performs a calculation and shows the result', async () => {
    const spy = vi
      .spyOn(api, 'calculate')
      .mockResolvedValue({ operation: 'add', a: 2, b: 3, result: 5 });
    const user = userEvent.setup();

    render(<Calculator />);
    await user.type(screen.getByLabelText('First value'), '2');
    await user.type(screen.getByLabelText('Second value'), '3');
    await user.click(screen.getByRole('button', { name: /calculate/i }));

    await waitFor(() => expect(screen.getByText('5')).toBeInTheDocument());
    expect(spy).toHaveBeenCalledWith('add', { a: 2, b: 3 });
  });

  it('shows client-side validation errors without calling the API', async () => {
    const spy = vi.spyOn(api, 'calculate');
    const user = userEvent.setup();

    render(<Calculator />);
    await user.click(screen.getByRole('button', { name: /calculate/i }));

    expect(await screen.findByText('First value is required')).toBeInTheDocument();
    expect(spy).not.toHaveBeenCalled();
  });

  it('surfaces backend errors such as division by zero', async () => {
    vi.spyOn(api, 'calculate').mockRejectedValue(
      new api.ApiError('division by zero is undefined', 422),
    );
    const user = userEvent.setup();

    render(<Calculator />);
    await user.click(screen.getByRole('button', { name: /divide/i }));
    await user.type(screen.getByLabelText('First value'), '1');
    await user.type(screen.getByLabelText('Second value'), '0');
    await user.click(screen.getByRole('button', { name: /calculate/i }));

    expect(await screen.findByText('division by zero is undefined')).toBeInTheDocument();
  });

  it('hides the second operand for unary operations like square root', async () => {
    const user = userEvent.setup();
    render(<Calculator />);

    await user.click(screen.getByRole('button', { name: /square root/i }));

    expect(screen.getByLabelText('First value')).toBeInTheDocument();
    expect(screen.queryByLabelText('Second value')).not.toBeInTheDocument();
  });

  it('labels the second field as percent for the percentage operation', async () => {
    const user = userEvent.setup();
    render(<Calculator />);

    await user.click(screen.getByRole('button', { name: /percent/i }));

    expect(screen.getByLabelText('Second value (percent)')).toBeInTheDocument();
  });
});

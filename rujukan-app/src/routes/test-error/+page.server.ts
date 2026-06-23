import { error } from '@sveltejs/kit';

export function load() {
  error(500, 'Internal Server Error: Simulasi kesalahan server.');
}

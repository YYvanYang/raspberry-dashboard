/** @type {import('@sveltejs/package').Config} */
export default {
	source: 'src/lib',
	dir: 'dist',
	emitTypes: true,
	exports: (filepath) => {
		if (filepath.endsWith('.svelte')) return true;
		return filepath === 'index.js' || filepath === 'index.ts';
	}
}; 
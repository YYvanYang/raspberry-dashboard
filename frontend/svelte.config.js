import adapter from '@sveltejs/adapter-static';

/** @type {import('@sveltejs/kit').Config} */
const config = {
    kit: {
        adapter: adapter(),
        alias: {
            $lib: './src/lib'
        }
    },
    compilerOptions: {
        runes: true
    }
};

export default config; 
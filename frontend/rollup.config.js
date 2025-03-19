import svelte from 'rollup-plugin-svelte';
import resolve from '@rollup/plugin-node-resolve';
import pkg from './package.json';
import css from 'rollup-plugin-css-only';
import terser from '@rollup/plugin-terser';
//import commonjs from '@rollup/plugin-commonjs';

const name = pkg.name
	.replace(/^(@\S+\/)?(svelte-)?(\S+)/, '$3')
	.replace(/^\w/, m => m.toUpperCase())
	.replace(/-\w/g, m => m[1].toUpperCase());

export default {
	input: 'src/index.js',
	output: {
		sourcemap: false,
		format: 'iife',
		name: 'app',
		file: '../web/js/scripts.min.js'
	},
	plugins: [
		css({ output: 'style.css' }),
		svelte(),
		//		commonjs(),
		resolve(),
		terser()
	]
};

import { useEffect, useRef } from 'react';

function App() {
	const goRef = useRef(null);
	const loadedRef = useRef(false)

	const loadWasm = async () => {
		const urlWasmEndPoint = "http://localhost:8080/game/game1/game1.wasm";
		try {
			const go = new Go();

			const res = await fetch(urlWasmEndPoint);
			const bytes = await res.arrayBuffer();
			if(!res.ok) {
				const message = bytes?.message || 'No response';
        		throw new Error(message, {cause: {response: res, bytes}});
			}

			const webAsm = await WebAssembly.instantiate(bytes, go.importObject);
			go.run(webAsm?.instance);

			return go;
		} catch(error) {
			console.log(error?.message || error);
			return null
		}
	}

	useEffect(() => {
		if (loadedRef.current) return;
		loadedRef.current = true;

		const go = loadWasm()
		goRef.current = go;

		// Cleanup if component unmounts
		return () => {
			goRef.current?.exit?.(0);
		};
	}, []);

	return null;
}

export default App;
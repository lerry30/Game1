import { useEffect, useRef } from 'react';

function App() {
  const goRef = useRef(null);
  const loadedRef = useRef(false)

  useEffect(() => {
	if(loadedRef.current) return;
	loadedRef.current = true;

    const go = new Go();
    goRef.current = go;

    fetch("http://localhost:8080/game/game1/game1.wasm")
      .then(res => {
        if (!res.ok) throw new Error(`Failed to fetch wasm: ${res.status}`);
        return res.arrayBuffer();
      })
      .then(bytes => WebAssembly.instantiate(bytes, go.importObject))
      .then(({ instance }) => {
        go.run(instance);
      })
      .catch(err => console.error("WASM load error:", err));

    // Cleanup if component unmounts
    return () => {
      goRef.current?.exit?.(0);
    };
  }, []);

  return null;
}

export default App;
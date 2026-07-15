import { Calculator } from './components/Calculator';
import './App.css';

export default function App() {
  return (
    <main className="app">
      <Calculator />
      <footer className="app__footer">
        Powered by a Go microservice · React + TypeScript
      </footer>
    </main>
  );
}

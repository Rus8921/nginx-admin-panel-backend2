import React from 'react';
import './App.css';
import Header from './components/Header/Header';
import WebsitesDashboard from "./routes/WebsitesDashboard/WebsitesDashboard";

function App() {
  return (
    <div className="App">
      <Header />
      <WebsitesDashboard />
    </div>
  );
}

export default App;

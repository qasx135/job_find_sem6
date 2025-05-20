import React from 'react';
import ReactDOM from 'react-dom/client';
import './App.css'; // или './index.css', если вы создали этот файл
import App from './App';
import { BrowserRouter } from 'react-router-dom';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <BrowserRouter>
      <App />
    </BrowserRouter>
  </React.StrictMode>
);
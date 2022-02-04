import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Home from "./pages/Home";
import 'antd/dist/antd.css';

const App = () => {
  return (
    <BrowserRouter>
      <Routes>
        {/*<Route path='/signin' element={<Login />} />*/}
        <Route path='/' element={<Home />} />
      </Routes>
    </BrowserRouter>
  );
};

export default App;

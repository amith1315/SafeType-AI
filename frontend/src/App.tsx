import React, { useState, useEffect } from 'react';
import './App.css'
import {BrowserRouter, Route, Routes} from "react-router-dom";
// import { LoginForm } from './components/login-form';
import MainPage from './comps/MainPage';
import { ThemeProvider } from "@/components/theme-provider"

function App() {

  // const [username, setUsername] = useState("");
  // const [password, setPassword] = useState("");

  return (
    <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
      <BrowserRouter>
        <div className="">
          <Routes>
            {/* { */}
            <Route path="/" element={<MainPage/>} /> 
            {/* <Route path="/login" element={<LoginForm user={user} setUser={setUser} setUsername={setUsername} setPassword={setPassword} username={username} password={password}/>} /> */}
            {/* } */}
            {/* <Route path="/checkout" element={<Checkout/>} cartData={cartData} check={check}/> */}
          </Routes>
        </div>
      </BrowserRouter>
    </ThemeProvider>
  )
}

export default App

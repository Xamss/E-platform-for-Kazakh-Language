import { BrowserRouter, Routes, Route } from 'react-router-dom';
import MyNavbar from './components/Navbar';
import LandingPage from './components/LandingPage/LandingPage';
import About from './components/About';
import Login from './components/Login';
import Signup from './components/Signup';
import Logout from './components/Logout';
import Map from './map/App'
// import Footer from "./components/Footer";
import './App.css';
import UserContext from './context/UserContext';
import React, { useState } from 'react';

// import LandingPage from "./components/LandingPage/LandingPage";
// for react-toastify
import 'react-toastify/dist/ReactToastify.css';
import { ToastContainer } from 'react-toastify';

function App() {
  const [user, setUser] = useState(localStorage.getItem('token') || null);

  return (
    <div>
      <BrowserRouter>
        <UserContext.Provider value={{ user, setUser }}>
          <MyNavbar />
          <ToastContainer />
          <Routes>
            <Route index path="/" element={<LandingPage />} />
            <Route path="/about" element={<About />} />
            <Route path="/login" element={<Login />} />
            <Route path={'/signup'} element={<Signup />} />
            <Route path={'/logout'} element={<Logout />} />
            <Route path={'/map'} element={<Map />} />
            <Route path="*" element={<LandingPage />} />
          </Routes>
          {/* <Footer /> */}
        </UserContext.Provider>
      </BrowserRouter>
    </div>
  );
}

export default App;

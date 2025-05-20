import { useState } from 'react';
import { Routes, Route, Navigate } from 'react-router-dom';
import { isAuthenticated } from './utils/auth';
import HomePage from './pages/HomePage';
import ProfilePage from './pages/ProfilePage';
import AuthPage from './pages/AuthPage';
import JobCreateForm from './components/Job/JobCreateForm';
import ResumeCreateForm from './components/Profile/ResumeCreateForm';
import JobDetails from './components/Job/JobDetails';
import Navbar from './components/Navbar';

const App = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(isAuthenticated());

  const handleLoginSuccess = () => {
    setIsLoggedIn(true);
  };

  const handleLogout = () => {
    localStorage.removeItem('user');
    setIsLoggedIn(false);
  };

  return (
    <>
      <Navbar isLoggedIn={isLoggedIn} onLogout={handleLogout} />
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/auth" element={
          isLoggedIn ? <Navigate to="/" /> : <AuthPage onLoginSuccess={handleLoginSuccess} />
        } />
        <Route path="/profile" element={
          isLoggedIn ? <ProfilePage /> : <Navigate to="/auth" />
        } />
        <Route path="/jobs/create" element={
          isLoggedIn ? <JobCreateForm /> : <Navigate to="/auth" />
        } />
        <Route path="/resumes/create" element={
          isLoggedIn ? <ResumeCreateForm /> : <Navigate to="/auth" />
        } />
        <Route path="/jobs/:id" element={<JobDetails />} />
      </Routes>
    </>
  );
};

export default App;
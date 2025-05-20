import { Link } from 'react-router-dom';
import { isAuthenticated } from '../utils/auth';

const Navbar = ({ isLoggedIn, onLogout }) => {
  return (
    <nav className="navbar">
      <div className="navbar-brand">
        <Link to="/">Job Portal</Link>
      </div>
      <div className="navbar-links">
        {isLoggedIn ? (
          <>
            <Link to="/profile">Profile</Link>
            <button onClick={onLogout}>Logout</button>
          </>
        ) : (
          <Link to="/auth">Login/Register</Link>
        )}
      </div>
    </nav>
  );
};

export default Navbar;
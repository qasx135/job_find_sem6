import { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { isAuthenticated, getAuthToken } from '../utils/auth';
import { getCurrentUser } from '../services/authService';
import { getUserJobs, getUserResumes } from '../services/jobService';
import ProfileInfo from '../components/Profile/ProfileInfo';
import VacancyList from '../components/Profile/VacancyList';
import ResumeList from '../components/Profile/ResumeList';

const ProfilePage = () => {
  const [activeTab, setActiveTab] = useState('profile');
  const [user, setUser] = useState(null);
  const [jobs, setJobs] = useState([]);
  const [resumes, setResumes] = useState([]);
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();

  useEffect(() => {
    if (!isAuthenticated()) {
      navigate('/auth');
      return;
    }

    const fetchData = async () => {
      try {
        const currentUser = getCurrentUser();
        setUser(currentUser);

        const token = getAuthToken();
        
        // Загружаем вакансии пользователя
        const userJobs = await getUserJobs(currentUser.id, token);
        setJobs(userJobs);

        // Загружаем резюме пользователя
        const userResumes = await getUserResumes(currentUser.id, token);
        setResumes(userResumes);

        setLoading(false);
      } catch (err) {
        console.error('Failed to fetch profile data', err);
        setLoading(false);
      }
    };

    fetchData();
  }, [navigate]);

  if (loading) return <div>Loading profile...</div>;
  if (!user) return <div>User not found</div>;

  return (
    <div className="profile-page">
      <div className="profile-tabs">
        <button 
          className={activeTab === 'profile' ? 'active' : ''}
          onClick={() => setActiveTab('profile')}
        >
          Profile
        </button>
        <button 
          className={activeTab === 'vacancies' ? 'active' : ''}
          onClick={() => setActiveTab('vacancies')}
        >
          My Vacancies
        </button>
        {isAuthenticated() && (
        <button onClick={() => navigate('/jobs/create')}>Create Vacancy</button>
        )}
        <button 
          className={activeTab === 'resumes' ? 'active' : ''}
          onClick={() => setActiveTab('resumes')}
        >
          My Resumes
        </button>
        {isAuthenticated() && (
        <button onClick={() => navigate('/resumes/create')}>Create Resume</button>
        )}
      </div>

      <div className="profile-content">
        {activeTab === 'profile' && <ProfileInfo user={user} />}
        {activeTab === 'vacancies' && <VacancyList jobs={jobs} />}
        {activeTab === 'resumes' && <ResumeList resumes={resumes} />}
      </div>
    </div>
  );
};

export default ProfilePage;
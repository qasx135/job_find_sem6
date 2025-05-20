import { useEffect, useState } from 'react';
import { getAllJobs } from '../services/jobService';
import JobCard from '../components/Job/JobCard';
import { isAuthenticated } from '../utils/auth';
import { useNavigate } from 'react-router-dom';

const HomePage = () => {
  const [jobs, setJobs] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const navigate = useNavigate();

  useEffect(() => {
    const fetchJobs = async () => {
      try {
        const data = await getAllJobs();
        setJobs(data);
        setLoading(false);
      } catch (err) {
        setError('Failed to load jobs');
        setLoading(false);
      }
    };

    fetchJobs();
  }, []);

  const handleJobClick = (jobId) => {
    navigate(`/jobs/${jobId}`);
  };

  if (loading) return <div>Loading...</div>;
  if (error) return <div>{error}</div>;

  return (
    <div className="home-page">
      <h1>Available Jobs</h1>
      <div className="job-list">
        {jobs.map((job) => (
          <JobCard 
            key={job.id} 
            job={job} 
            onClick={() => handleJobClick(job.id)} 
          />
        ))}
      </div>
    </div>
  );
};

export default HomePage;
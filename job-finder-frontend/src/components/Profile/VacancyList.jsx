import { useNavigate } from 'react-router-dom';
import JobCard from '../Job/JobCard';

const VacancyList = ({ jobs }) => {
  const navigate = useNavigate();

  if (jobs.length === 0) {
    return <p>No vacancies found</p>;
  }
  const handleJobClick = (jobId) => {
    navigate(`/jobs/${jobId}`);
  };

  return (
    <div className="vacancy-list">
      <h3>My Vacancies</h3>
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

export default VacancyList;
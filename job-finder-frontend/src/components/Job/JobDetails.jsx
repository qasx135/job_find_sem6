import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { getJobDetails } from '../../services/jobService';

const JobDetails = () => {
  const { id } = useParams();
  const [job, setJob] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    const fetchJob = async () => {
      try {
        const data = await getJobDetails(id);
        setJob(data);
        setLoading(false);
      } catch (err) {
        setError('Failed to load job details');
        setLoading(false);
      }
    };

    fetchJob();
  }, [id]);

  if (loading) return <div>Loading...</div>;
  if (error) return <div>{error}</div>;
  if (!job) return <div>Job not found</div>;

  return (
    <div className="job-details">
      <h2>header: {job.header}</h2>
      <p>salary: {job.salary}</p>
      <h3>experience: {job.experience}</h3>
      <p>employment: {job.employment}</p>
      <p>schedule: {job.schedule}</p>
      <p>work_format: {job.work_format}</p>
      <p>working_hours: {job.working_hours}</p>
      <p>description: {job.description}</p>
    </div>
  );
};

export default JobDetails;
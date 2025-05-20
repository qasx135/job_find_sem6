const JobCard = ({ job, onClick }) => {
  console.log(job)
  return (
    <div className="job-card" onClick={onClick}>
      <h3>Header: {job.header}</h3>
      <p>Experience: {job.experience}</p>
      <p>Employment: {job.employment}</p>
      <p>Salary: {job.salary}</p>
    </div>
  );
};

export default JobCard;
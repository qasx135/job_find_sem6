import JobCard from '../Job/ResumeCard';
import ResumeCard from '../Job/ResumeCard';
const ResumeList = ({ resumes }) => {
  if (resumes.length === 0) {
    return <p>No resumes found</p>;
  }

  return (
    <div className="resume-list">
      <h3>My Resumes</h3>
      <div className="resume-list">
        {resumes.map((resume) => (
          <ResumeCard
            key={resume.id}
            resume={resume}
          />
        ))}
      </div>
    </div>
  );
};

export default ResumeList;
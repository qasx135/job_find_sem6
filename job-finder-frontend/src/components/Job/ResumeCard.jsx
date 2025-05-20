const ResumeCard = ({ resume }) => {
  return (
    <div className="resume-card">
      <h3>About: {resume.about}</h3>
      <p>Experience: {resume.experience}</p>
    </div>
  );
};

export default ResumeCard;
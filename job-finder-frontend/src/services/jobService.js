// ЗАМЕНИТЕ BASE_URL на адрес вашего бэкенда
const BASE_URL = '/api';

export const getAllJobs = async () => {
  const response = await fetch(`${BASE_URL}/all-jobs-home`);
  return response.json();
};

export const createJob = async (jobData, token) => {
  const response = await fetch(`${BASE_URL}/new-job`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${token}`,
    },
    body: JSON.stringify(jobData),
  });
  return response.json();
};

export const getUserJobs = async (userId, token) => {
  const response = await fetch(`${BASE_URL}/all-jobs`, {
    headers: {
      'Authorization': `Bearer ${token}`,
    },
  });
  return response.json();
};

export const createResume = async (resumeData, token) => {
  const response = await fetch(`${BASE_URL}/new-resume`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${token}`,
    },
    body: JSON.stringify(resumeData),
  });
  return response.json();
};

export const getUserResumes = async (userId, token) => {
  const response = await fetch(`${BASE_URL}/all-resumes`, {
    headers: {
      'Authorization': `Bearer ${token}`,
    },
  });
  return response.json();
};

export const getJobDetails = async (jobId) => {
  const response = await fetch(`${BASE_URL}/jobs/${jobId}`);
  return response.json();
};
// src/components/JobCreateForm.jsx

import React from 'react';
import { useForm } from 'react-hook-form';
import { useNavigate } from 'react-router-dom';
import { createResume } from '../../services/jobService';
import { getAuthToken } from '../../utils/auth';

const ResumeCreateForm = () => {
  const navigate = useNavigate();
  const token = getAuthToken();

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm();

  const onSubmit = async (data) => {
    console.log('Form data:', data);
    console.log('Token:', token);
    try {
      await createResume(data, token);
      navigate('/');
    } catch (err) {
      console.error('Failed to create job', err);
    }
  };

  return (
    <div className="resume-form">
      <h2>Create New Resume</h2>

      <form onSubmit={handleSubmit(onSubmit)}>
        {/* Header */}
        <div>
          <label>About</label>
          <textarea
            type="text"
            {...register('about', { required: 'About is required' })}
            placeholder="Description"
          />
          {errors.header && <span className="error">{errors.header.message}</span>}
        </div>

        {/* Experience */}
        <div>
          <label>Experience</label>
          <input
            type="text"
            {...register('experience', { required: 'Experience is required' })}
            placeholder="e.g. 3 years"
          />
          {errors.experience && <span className="error">{errors.experience.message}</span>}
        </div>

        <button type="submit">Create Resume</button>
      </form>
    </div>
  );
};

export default ResumeCreateForm;
// src/components/JobCreateForm.jsx

import React from 'react';
import { useForm } from 'react-hook-form';
import { useNavigate } from 'react-router-dom';
import { createJob } from '../../services/jobService';
import { getAuthToken } from '../../utils/auth';

const JobCreateForm = () => {
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
      await createJob(data, token);
      navigate('/');
    } catch (err) {
      console.error('Failed to create job', err);
    }
  };

  return (
    <div className="job-form">
      <h2>Create New Job</h2>

      <form onSubmit={handleSubmit(onSubmit)}>
        {/* Header */}
        <div>
          <label>Header</label>
          <input
            type="text"
            {...register('header', { required: 'Header is required' })}
            placeholder="Job title"
          />
          {errors.header && <span className="error">{errors.header.message}</span>}
        </div>
        {/* Salary */}
        <div>
          <label>Salary</label>
          <input
            type="text"
            {...register('salary', { required: 'Salary is required' })}
            placeholder="Salary"
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

        {/* Employment */}
        <div>
          <label>Employment</label>
          <input
            type="text"
            {...register('employment', { required: 'Employment is required' })}
            placeholder="Full-time / Part-time"
          />
          {errors.employment && <span className="error">{errors.employment.message}</span>}
        </div>

        {/* Schedule */}
        <div>
          <label>Schedule</label>
          <input
            type="text"
            {...register('schedule', { required: 'Schedule is required' })}
            placeholder="e.g. Remote, Office"
          />
          {errors.schedule && <span className="error">{errors.schedule.message}</span>}
        </div>

        {/* Work Format */}
        <div>
          <label>Work Format</label>
          <input
            type="text"
            {...register('work_format', { required: 'Work format is required' })}
            placeholder="e.g. Full-time, Flexible"
          />
          {errors.work_format && <span className="error">{errors.work_format.message}</span>}
        </div>

        {/* Working Hours */}
        <div>
          <label>Working Hours</label>
          <input
            type="text"
            {...register('working_hours', { required: 'Working hours are required' })}
            placeholder="e.g. 9-5"
          />
          {errors.working_hours && <span className="error">{errors.working_hours.message}</span>}
        </div>

        {/* Description */}
        <div>
          <label>Description</label>
          <textarea
            {...register('description', { required: 'Description is required' })}
            placeholder="Job description"
          />
          {errors.description && <span className="error">{errors.description.message}</span>}
        </div>

        <button type="submit">Create Job</button>
      </form>
    </div>
  );
};

export default JobCreateForm;
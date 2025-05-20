const ProfileInfo = ({ user }) => {
  return (
    <div className="profile-info">
      <h2>{user.name}</h2>
      <p>Email: {user.email}</p>
      {/* Добавьте другие поля профиля по мере необходимости */}
    </div>
  );
};

export default ProfileInfo;
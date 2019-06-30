import React from 'react';

function ConfirmWidget({ handleSubmit, options }) {
  return (
    <>
      <button onClick={handleSubmit()}>Update</button>
    </>
  );
}

export default ConfirmWidget;

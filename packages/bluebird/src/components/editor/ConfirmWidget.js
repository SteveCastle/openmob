import React from 'react';

function ConfirmWidget({ handleSubmit, options }) {
  return (
    <div
      onClick={e => {
        e.stopPropagation();
      }}
    >
      <button onClick={handleSubmit()}>Update</button>
    </div>
  );
}

export default ConfirmWidget;

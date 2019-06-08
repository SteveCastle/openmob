import React from 'react';

function PageEditor({ children }) {
  return (
    <div
      style={{
        minHeight: 'min-content',
        display: 'flex',
        flexWrap: 'wrap',
        marginBottom: '20px',
        width: '100%',
      }}
    >
      {children}
    </div>
  );
}

export default PageEditor;

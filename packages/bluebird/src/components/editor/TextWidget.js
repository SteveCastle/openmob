import React, { useState } from 'react';
import PropTypes from 'prop-types';

function TextWidget({ handleSubmit, title, initValue }) {
  const [value, setValue] = useState(false);

  return (
    <div
      onClick={e => {
        e.stopPropagation();
      }}
    >
      <span>{title}</span>
      <input
        type="text"
        onChange={e => setValue(e.target.value)}
        value={value || initValue}
      />
      <button onClick={handleSubmit(value)}>Update</button>
    </div>
  );
}

TextWidget.propTypes = {
  handleSubmit: PropTypes.func,
  title: PropTypes.string,
  initValue: PropTypes.string
};

export default TextWidget;

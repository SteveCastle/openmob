import React, { useState } from 'react';
import PropTypes from 'prop-types';

function TextWidget({ handleSubmit, title, initValue }) {
  const [value, setValue] = useState('');
  const [fresh, setFresh] = useState(true);

  return (
    <div
      onClick={e => {
        e.stopPropagation();
      }}
    >
      <span>{title}</span>
      <input
        type="text"
        onChange={e => {
          setValue(e.target.value);
          setFresh(false);
        }}
        value={initValue && fresh ? initValue : value}
      />
      <button onClick={handleSubmit(value)}>Update</button>
    </div>
  );
}

TextWidget.propTypes = {
  handleSubmit: PropTypes.func,
  title: PropTypes.string,
  initValue: PropTypes.oneOfType([PropTypes.number, PropTypes.string])
};

export default TextWidget;

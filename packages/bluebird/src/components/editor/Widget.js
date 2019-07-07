import React, { useState } from 'react';
import PropTypes from 'prop-types';

function Widget({ handleSubmit, initValue, options }) {
  const [value, setValue] = useState(false);

  return (
    <div
      onClick={e => {
        e.stopPropagation();
      }}
    >
      {options && (
        <select
          onChange={e => setValue(e.target.value)}
          value={value || initValue}
        >
          {options.map(item => (
            <option value={item.ID} key={item.ID}>
              {item.Title}
            </option>
          ))}
        </select>
      )}
      <button onClick={handleSubmit(value)}>Update</button>
    </div>
  );
}

Widget.propTypes = {
  handleSubmit: PropTypes.func,
  initValue: PropTypes.oneOfType([PropTypes.number, PropTypes.string]),
  options: PropTypes.arrayOf(
    PropTypes.shape({
      ID: PropTypes.oneOfType([PropTypes.string, PropTypes.number]),
      Title: PropTypes.string
    })
  )
};

export default Widget;

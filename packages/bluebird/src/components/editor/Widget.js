import React, { useState } from 'react';
import PropTypes from 'prop-types';

function Widget({ handleSubmit, options }) {
  const [value, setValue] = useState(false);

  return (
    <div>
      {options && (
        <select onChange={e => setValue(e.target.value)}>
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
  options: PropTypes.arrayOf(
    PropTypes.shape({
      ID: PropTypes.string,
      Title: PropTypes.string
    })
  )
};

export default Widget;

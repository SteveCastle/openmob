import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import { space, fontSize, color, borders, borderRadius } from 'styled-system';

const StyledControl = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
  ${fontSize}
  ${color}
`;

function Control({ onClick, children }) {
  return (
    <StyledControl border={0} color={'gray.1'}>
      {children}
    </StyledControl>
  );
}

Control.propTypes = {
  onClick: PropTypes.func
};
Control.defaultProps = {};

export default Control;

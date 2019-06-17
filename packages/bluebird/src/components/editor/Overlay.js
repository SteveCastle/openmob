import React, { useState } from 'react';

import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  space,
  fontSize,
  color,
  borders,
  borderRadius,
  display
} from 'styled-system';

const StyledOverlay = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
  ${fontSize}
  ${color}
  ${display}
  align-items: center;
  justify-content: center;
  position: absolute;
  box-sizing: border-box;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  opacity: ${({ locked }) => (locked ? 0.9 : 0)};
  transition: opacity .2s ease-out;
  z-index: ${({ locked }) => (locked ? 999 : 998)};;
  & :hover {
    opacity: ${({ locked }) => (locked ? 0.9 : 0.8)};

  }
`;

function Overlay({ children, locked, onClick }) {
  return (
    <StyledOverlay
      border={0}
      bg={'gray.9'}
      locked={locked}
      onClick={onClick}
      display="flex"
    >
      {children}
    </StyledOverlay>
  );
}

Overlay.propTypes = {
  children: PropTypes.oneOfType([
    PropTypes.arrayOf(PropTypes.node),
    PropTypes.node
  ]),
  locked: PropTypes.bool,
  onClick: PropTypes.func
};
Overlay.defaultProps = {};

export default Overlay;

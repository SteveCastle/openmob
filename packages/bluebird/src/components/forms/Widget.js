import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  space,
  width,
  fontSize,
  color,
  borders,
  borderRadius
} from 'styled-system';

const StyledWidget = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${color}
`;

function Widget({ children, block }) {
  return (
    <StyledWidget border={0} p={3} fontSize={1} width={block ? 1 : null}>
      {children}
    </StyledWidget>
  );
}

Widget.propTypes = {
  children: PropTypes.element,
  state: PropTypes.oneOf(['ready', 'valid', 'error'])
};
export default Widget;

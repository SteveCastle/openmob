import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  space,
  width,
  fontSize,
  color,
  borders,
  borderRadius,
  display,
  flexWrap,
  maxWidth
} from 'styled-system';

const StyledContainer = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${color}
  ${display}
  ${flexWrap}
  ${maxWidth}
  height: 100%
`;

function Container({ dark, children, tracing }) {
  return (
    <StyledContainer
      display={'flex'}
      border={0}
      flexWrap="wrap"
      m="auto"
      p={0}
      width={1}
      maxWidth={7}
      bg={tracing ? `violet.${tracing}` : null}
      fontSize={1}
    >
      {children}
    </StyledContainer>
  );
}

Container.propTypes = {
  children: PropTypes.oneOfType([
    PropTypes.arrayOf(PropTypes.node),
    PropTypes.node
  ]),
  dark: PropTypes.bool,
  tracing: PropTypes.number
};
Container.defaultProps = {};

export default Container;

import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  space,
  width,
  height,
  fontSize,
  color,
  borders,
  borderRadius,
  display,
  flexWrap
} from 'styled-system';

const StyledContent = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${height}
  ${fontSize}
  ${flexWrap}
  ${color}
  ${display}
`;

function Content({ children }) {
  return (
    <StyledContent
      display={'flex'}
      border={0}
      flexWrap="wrap"
      width={1}
      height={5}
      fontSize={1}
    >
      {children}
    </StyledContent>
  );
}

Content.propTypes = {
  children: PropTypes.arrayOf(PropTypes.element),
};
Content.defaultProps = {};

export default Content;

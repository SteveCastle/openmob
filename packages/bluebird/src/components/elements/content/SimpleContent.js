import React from 'react';
import styled from 'styled-components';
import PropTypes from 'prop-types';

import {
  space,
  width,
  fontSize,
  color,
  borders,
  borderRadius,
  display,
  alignItems,
  justifyContent
} from 'styled-system';

const StyledContent = styled.section`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${alignItems}
  ${justifyContent}
  ${color}
  ${display}
  min-height: 400px;
  `;
const Title = styled.h1`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${alignItems}
  ${justifyContent}
  ${color}
  ${display}
  `;

function SimpleContent({ title }) {
  return (
    <StyledContent
      display={'flex'}
      justifyContent="center"
      alignItems="center"
      bg="cyan.3"
    >
      <Title color="gray.0" size={1}>
        {title}
      </Title>
    </StyledContent>
  );
}

SimpleContent.propTypes = {
  title: PropTypes.string.isRequired
};
SimpleContent.defaultProps = { title: 'Simple Content' };

export default SimpleContent;

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

const StyledFooter = styled.section`
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

function SimpleFooter({ title }) {
  return (
    <StyledFooter
      display={'flex'}
      justifyContent="center"
      alignItems="flex-end"
      bg="blue.8"
    >
      <Title color="gray.0" size={1}>
        {title}
      </Title>
    </StyledFooter>
  );
}

SimpleFooter.propTypes = {
  title: PropTypes.shape({
    StringValue: PropTypes.string.isRequired
  })
};
SimpleFooter.defaultProps = {};

export default SimpleFooter;

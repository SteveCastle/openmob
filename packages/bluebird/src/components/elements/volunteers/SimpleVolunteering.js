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
  justifyContent,
} from 'styled-system';

const StyledVolunteering = styled.section`
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

function SimpleVolunteering({ title }) {
  return (
    <StyledVolunteering
      display={'flex'}
      justifyContent="center"
      alignItems="center"
      bg="red.4"
    >
      <Title color="gray.0" size={1}>
        {title}
      </Title>
    </StyledVolunteering>
  );
}

SimpleVolunteering.propTypes = {
  title: PropTypes.string.isRequired,
};
SimpleVolunteering.defaultProps = { title: 'Simple Volunteering' };

export default SimpleVolunteering;

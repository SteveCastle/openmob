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

const StyledHero = styled.section`
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
  background-image: url("${props => props.image}");
  background-size: cover;
  background-position: center;
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

function ImageHero({
  title,
  image = 'https://punknaturalism.com/static/4f45e3f86b23864cb0c71a1f3353b460/a5547/C9up3kLgxni.jpg',
}) {
  return (
    <StyledHero
      display={'flex'}
      justifyContent="center"
      alignItems="center"
      width={1}
      image={image}
    >
      <Title color="gray.0" size={4}>
        {title}
      </Title>
    </StyledHero>
  );
}

ImageHero.propTypes = {
  title: PropTypes.string.isRequired,
  image: PropTypes.string.isRequired,
};
ImageHero.defaultProps = {};

export default ImageHero;

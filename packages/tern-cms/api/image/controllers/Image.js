'use strict';

/**
 * Image.js controller
 *
 * @description: A set of functions called "actions" for managing `Image`.
 */

module.exports = {

  /**
   * Retrieve image records.
   *
   * @return {Object|Array}
   */

  find: async (ctx) => {
    if (ctx.query._q) {
      return strapi.services.image.search(ctx.query);
    } else {
      return strapi.services.image.fetchAll(ctx.query);
    }
  },

  /**
   * Retrieve a image record.
   *
   * @return {Object}
   */

  findOne: async (ctx) => {
    return strapi.services.image.fetch(ctx.params);
  },

  /**
   * Count image records.
   *
   * @return {Number}
   */

  count: async (ctx) => {
    return strapi.services.image.count(ctx.query);
  },

  /**
   * Create a/an image record.
   *
   * @return {Object}
   */

  create: async (ctx) => {
    return strapi.services.image.add(ctx.request.body);
  },

  /**
   * Update a/an image record.
   *
   * @return {Object}
   */

  update: async (ctx, next) => {
    return strapi.services.image.edit(ctx.params, ctx.request.body) ;
  },

  /**
   * Destroy a/an image record.
   *
   * @return {Object}
   */

  destroy: async (ctx, next) => {
    return strapi.services.image.remove(ctx.params);
  },

  /**
   * Add relation to a/an image record.
   *
   * @return {Object}
   */

  createRelation: async (ctx, next) => {
    return strapi.services.image.addRelation(ctx.params, ctx.request.body);
  },

  /**
   * Update relation to a/an image record.
   *
   * @return {Object}
   */

  updateRelation: async (ctx, next) => {
    return strapi.services.image.editRelation(ctx.params, ctx.request.body);
  },

  /**
   * Destroy relation to a/an image record.
   *
   * @return {Object}
   */

  destroyRelation: async (ctx, next) => {
    return strapi.services.image.removeRelation(ctx.params, ctx.request.body);
  }
};

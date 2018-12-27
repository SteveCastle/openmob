'use strict';

/**
 * Cause.js controller
 *
 * @description: A set of functions called "actions" for managing `Cause`.
 */

module.exports = {

  /**
   * Retrieve cause records.
   *
   * @return {Object|Array}
   */

  find: async (ctx) => {
    if (ctx.query._q) {
      return strapi.services.cause.search(ctx.query);
    } else {
      return strapi.services.cause.fetchAll(ctx.query);
    }
  },

  /**
   * Retrieve a cause record.
   *
   * @return {Object}
   */

  findOne: async (ctx) => {
    return strapi.services.cause.fetch(ctx.params);
  },

  /**
   * Count cause records.
   *
   * @return {Number}
   */

  count: async (ctx) => {
    return strapi.services.cause.count(ctx.query);
  },

  /**
   * Create a/an cause record.
   *
   * @return {Object}
   */

  create: async (ctx) => {
    return strapi.services.cause.add(ctx.request.body);
  },

  /**
   * Update a/an cause record.
   *
   * @return {Object}
   */

  update: async (ctx, next) => {
    return strapi.services.cause.edit(ctx.params, ctx.request.body) ;
  },

  /**
   * Destroy a/an cause record.
   *
   * @return {Object}
   */

  destroy: async (ctx, next) => {
    return strapi.services.cause.remove(ctx.params);
  },

  /**
   * Add relation to a/an cause record.
   *
   * @return {Object}
   */

  createRelation: async (ctx, next) => {
    return strapi.services.cause.addRelation(ctx.params, ctx.request.body);
  },

  /**
   * Update relation to a/an cause record.
   *
   * @return {Object}
   */

  updateRelation: async (ctx, next) => {
    return strapi.services.cause.editRelation(ctx.params, ctx.request.body);
  },

  /**
   * Destroy relation to a/an cause record.
   *
   * @return {Object}
   */

  destroyRelation: async (ctx, next) => {
    return strapi.services.cause.removeRelation(ctx.params, ctx.request.body);
  }
};

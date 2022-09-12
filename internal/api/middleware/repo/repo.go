// Copyright 2021 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package repo

import (
	"net/http"
	"strconv"

	"github.com/harness/gitness/internal/api/render"
	"github.com/harness/gitness/internal/api/request"
	"github.com/harness/gitness/internal/store"
	"github.com/harness/gitness/types"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

/*
 * Required returns an http.HandlerFunc middleware that resolves the
 * repository using the fqrn from the request and injects the repo into the request.
 * In case the fqrn isn't found or the repository doesn't exist an error is rendered.
 */
func Required(repos store.RepoStore) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := hlog.FromRequest(r)

			ref, err := request.GetRepoRef(r)
			if err != nil {
				render.BadRequest(w)
				return
			}

			var repo *types.Repository

			// check if ref is repoId - ASSUMPTION: digit only is no valid repo name
			id, err := strconv.ParseInt(ref, 10, 64)
			if err == nil {
				repo, err = repos.Find(ctx, id)
			} else {
				repo, err = repos.FindByPath(ctx, ref)
			}

			if err != nil {
				log.Debug().Err(err).Msgf("Failed to get repo using ref '%s'.", ref)

				render.UserfiedErrorOrInternal(w, err)
				return
			}

			// Update the logging context and inject repo in context
			log.UpdateContext(func(c zerolog.Context) zerolog.Context {
				return c.Int64("repo_id", repo.ID).Str("repo_path", repo.Path)
			})

			next.ServeHTTP(w, r.WithContext(
				request.WithRepo(ctx, repo),
			))
		})
	}
}
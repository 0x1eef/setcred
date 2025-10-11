#include "control.h"
#include <errno.h>
#include <fcntl.h>
#include <stdio.h>
#include <unistd.h>

static int
set_feature(hbsdctrl_ctx_t*, const char*, const char*, int);

int
enable_feature(hbsdctrl_ctx_t* ctx, const char* name, const char* path)
{
  return set_feature(ctx, name, path, 1);
}

int
disable_feature(hbsdctrl_ctx_t* ctx, const char* name, const char* path)
{
  return set_feature(ctx, name, path, 2);
}

int
sysdef_feature(hbsdctrl_ctx_t* ctx, const char* name, const char* path)
{
  return set_feature(ctx, name, path, 3);
}

static int
set_feature(hbsdctrl_ctx_t* ctx, const char* name, const char* path, int state)
{
  hbsdctrl_feature_t* feature = NULL;
  int fd, res;

  res = -1;
  errno = 0;
  feature = NULL;
  fd = -1;
  if (ctx == NULL)
    goto done;
  feature = hbsdctrl_ctx_find_feature_by_name(ctx, name);
  if (feature == NULL)
    goto done;
  fd = open(path, O_PATH);
  if (fd == -1)
    goto done;
  if (feature->hf_apply(ctx, feature, &fd, &state) == RES_FAIL)
    goto done;
  res = 0;
  goto done;
done:
  if (fd != -1)
    close(fd);
  if (errno != 0)
    return errno;
  return res;
}

int
feature_status(hbsdctrl_ctx_t* ctx,
               const char* name,
               const char* path,
               char** status)
{
  hbsdctrl_feature_t* feature;
  hbsdctrl_feature_state_t state;
  int res = 0, fd = -1;

  errno = 0;
  if (ctx == NULL) {
    res = -1;
    goto done;
  }
  if ((fd = open(path, O_PATH)) == -1) {
    res = -1;
    goto done;
  }
  if ((feature = hbsdctrl_ctx_find_feature_by_name(ctx, name)) == NULL) {
    res = -1;
    goto done;
  }
  if (feature->hf_get(ctx, feature, &fd, &state) == RES_FAIL) {
    res = -1;
    goto done;
  }
  *status = (char*)hbsdctrl_feature_state_to_string(&state);
  goto done;
done:
  if (fd != -1)
    close(fd);
  if (errno != 0)
    return errno;
  return res;
}
import { useMemo, useState } from "react";

import {
  Box,
  Button,
  Chip,
  Divider,
  TextField,
  Tooltip,
  Typography,
} from "@mui/material";

import { enqueueSnackbar } from "notistack";
import { ConfigField, DeleteCheckModal, Dropdown, Multiselect } from "../..";
import {
  ChecksQuery,
  useDeleteCheckMutation,
  useUpdateCheckMutation,
  useValidateCheckMutation,
} from "../../../graph";

type props = {
  check: ChecksQuery["checks"][0];
  handleRefetch: () => void;
  visible: boolean;
};

export default function EditCheck({ check, visible, handleRefetch }: props) {
  const [expanded, setExpanded] = useState(false);

  const [config, setConfig] = useState<{
    [key: string]: number | boolean | string;
  }>(JSON.parse(check.config));
  const configChanged = useMemo(
    () => JSON.stringify(config) != check.config,
    [config, check.config]
  );

  const [editableFields, setEditableFields] = useState<string[]>(
    check.editable_fields
  );
  const editableFieldsChanged = useMemo(() => {
    return (
      [...editableFields].sort().join(",") !=
      [...check.editable_fields].sort().join(",")
    );
  }, [editableFields, check.editable_fields]);

  const [name, setName] = useState<string>(check.name);
  const nameChanged = useMemo(() => name != check.name, [name, check.name]);

  const [weight, setWeight] = useState<number>(check.weight);
  const weightChanged = useMemo(
    () => weight != check.weight,
    [weight, check.weight]
  );

  const [open, setOpen] = useState(false);

  const [validated, setValidated] = useState(true);
  const [validationError, setValidationError] = useState<string | undefined>();

  const [updateCheckMutation] = useUpdateCheckMutation({
    onCompleted: () => {
      enqueueSnackbar("Check updated successfully", { variant: "success" });
      handleRefetch();
    },
    onError: (error) => {
      enqueueSnackbar(error.message, { variant: "error" });
      console.error(error);
    },
  });

  const [DeleteCheckMutation] = useDeleteCheckMutation({
    onCompleted: () => {
      enqueueSnackbar("Check deleted successfully", { variant: "success" });
      handleRefetch();
    },
    onError: (error) => {
      enqueueSnackbar(error.message, { variant: "error" });
      console.error(error);
    },
  });

  const [validateCheckMutation] = useValidateCheckMutation({
    onCompleted: () => {
      setValidated(true);
      setValidationError(undefined);
      enqueueSnackbar("Check successfully validated", { variant: "success" });
    },
    onError: (error) => {
      enqueueSnackbar("Check failed to be validate", { variant: "error" });
      setValidationError(error.message);
    },
  });

  const handleConfigChange = (
    key: string,
    value: string | number | boolean
  ) => {
    setValidated(false);
    setConfig({ ...config, [key]: value });
  };

  const handleSave = () => {
    updateCheckMutation({
      variables: {
        id: check.id,
        name: nameChanged ? name : undefined,
        weight: weightChanged ? weight : undefined,
        config: configChanged ? JSON.stringify(config) : undefined,
        editable_fields: editableFieldsChanged ? editableFields : undefined,
      },
    });
  };

  const handleDelete = () => {
    DeleteCheckMutation({
      variables: {
        id: check.id,
      },
    });
  };

  return (
    <Dropdown
      modal={
        <DeleteCheckModal
          check={check.name}
          open={open}
          setOpen={setOpen}
          handleDelete={handleDelete}
        />
      }
      title={
        <>
          {expanded ? (
            <TextField
              label='Name'
              value={name}
              onClick={(e) => {
                e.stopPropagation();
              }}
              onChange={(e) => {
                setName(e.target.value);
              }}
              sx={{ marginRight: "24px" }}
              size='small'
            />
          ) : (
            <Typography variant='h6' component='div' marginRight='24px'>
              {check.name}
            </Typography>
          )}
          <Typography
            variant='subtitle1'
            color='textSecondary'
            component='div'
            marginRight='24px'
          >
            {check.source.name}
          </Typography>
          {expanded ? (
            <TextField
              label='Weight'
              type='number'
              value={weight}
              onClick={(e) => {
                e.stopPropagation();
              }}
              onChange={(e) => {
                setWeight(parseInt(e.target.value));
              }}
              sx={{ marginRight: "24px", width: "100px" }}
              size='small'
            />
          ) : (
            <Chip size='small' label={`weight:${weight}`} />
          )}
        </>
      }
      expandableButtons={[
        <Button
          variant='contained'
          onClick={() => {
            setOpen(true);
          }}
          color='error'
        >
          Delete
        </Button>,
        <Button
          variant='contained'
          onClick={(event) => {
            event.stopPropagation();
            validateCheckMutation({
              variables: {
                source: check.source.name,
                config: JSON.stringify(config),
              },
            });
          }}
        >
          Validate
        </Button>,
      ]}
      toggleButton={
        <Tooltip
          title={
            validated
              ? ""
              : "Configuration must first be validated before saving"
          }
        >
          <Box sx={{ display: "contents" }}>
            <Button
              variant='contained'
              color='success'
              disabled={!validated}
              onClick={(e) => {
                if (!expanded) {
                  e.stopPropagation();
                }

                handleSave();
              }}
            >
              Save
            </Button>
          </Box>
        </Tooltip>
      }
      expanded={expanded}
      setExpanded={setExpanded}
      visible={visible}
      toggleButtonVisible={
        configChanged || nameChanged || editableFieldsChanged || weightChanged
      }
    >
      <Box
        sx={{
          display: "flex",
          gap: "16px",
          flexWrap: "wrap",
          justifyContent: "center",
        }}
      >
        {Object.entries(check.source.schema).map(([_, field]) => (
          <ConfigField
            key={field.name}
            handleInputChange={handleConfigChange}
            fieldName={field.name}
            fieldType={field.type}
            checkConfig={config}
            defaultValue={field.default ?? undefined}
            enumValues={field.enum ?? undefined}
          />
        ))}
      </Box>
      {validationError !== undefined && (
        <>
          <Divider sx={{ margin: "16px 20% 20px 20%" }} />
          <TextField
            fullWidth
            label='Validation Error'
            multiline
            error
            value={validationError}
            InputProps={{
              readOnly: true,
            }}
          />
        </>
      )}
      <Divider sx={{ margin: "16px 20% 20px 20%" }} />
      <Multiselect
        label='Set User Editable Fields'
        placeholder='Select fields'
        options={Object.keys(config)}
        selected={editableFields}
        setSelected={setEditableFields}
      />
    </Dropdown>
  );
}

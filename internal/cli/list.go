package cli

import (
	"encoding/json"
	"fmt"
	"slices"

	"github.com/theshizik3/picoconv/internal/converters"

	"github.com/spf13/cobra"
)

var flat, json_ bool

var listCmd = &cobra.Command{
	Use:   "list [flags] source target",
	Short: "Displays all available formats",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return list(flat, json_, src, trg)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	f := listCmd.Flags()
	f.BoolVar(&flat, "flat", false, "outputs only types")
	f.BoolVar(&json_, "json", false, "outputs as JSON")
	f.StringVar(&src, "src", "", "output only converters of this type")
	f.StringVar(&trg, "trg", "", "outputs only converters to this type")
}

func list(flat, json_ bool, src, trg string) error {
	reg, err := converters.BuildRegistry()
	if err != nil {
		return err
	}

	l := reg.SupportedTypes()

	if src != "" {
		l = slices.DeleteFunc(l, func(p [2]string) bool {
			return p[0] != src
		})
	}

	if trg != "" {
		l = slices.DeleteFunc(l, func(p [2]string) bool {
			return p[1] != trg
		})
	}

	switch {
	case flat && json_:
		l := flatten(l)
		return json_output(l)

	case flat && !json_:
		l := flatten(l)
		for _, el := range l {
			fmt.Printf("%s ", el)
		}
		fmt.Println()

	case !flat && json_:
		m := make(map[string][]string)
		for _, el := range l {
			if v, ok := m[el[0]]; !ok {
				m[el[0]] = []string{el[1]}
			} else {
				m[el[0]] = append(v, el[1])
			}
		}
		return json_output(m)

	case !flat && !json_:
		m := make(map[string][]string)
		for _, el := range l {
			if v, ok := m[el[0]]; !ok {
				m[el[0]] = []string{el[1]}
			} else {
				m[el[0]] = append(v, el[1])
			}
		}
		for k := range m {
			fmt.Printf("%s -> %v\n", k, m[k])
		}
	}
	return nil
}

func json_output(v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

func flatten(l [][2]string) []string {
	new_l := []string{}
	check := make(map[string]bool)
	for _, k := range l {
		for _, p := range k {
			if _, ok := check[p]; !ok {
				check[p] = true
				new_l = append(new_l, p)
			}
		}
	}
	return new_l
}
